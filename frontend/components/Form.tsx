import { FormEvent, useState } from "react";

export enum FormState {
  IDLE,
  SUBMITTING,
  SUCCESS,
  ERROR,
}
const Form = () => {
  const [formState, setFormState] = useState<FormState>(FormState.IDLE);

  const [file, setFile] = useState<File>();

  const [downloadUri, setDownUri] = useState<string>();

  const [error, setError] = useState("");

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData();

    if (!file) {
      setFormState(FormState.ERROR);
      setError("no file selected!");
      return;
    }
    formData.append("file", file);

    try {
      setFormState(FormState.SUBMITTING);
      setError("");
      setDownUri("");
      const response = await fetch("/ocr", {
        method: "post",
        body: formData,
      });
      const data = response.url;
      setDownUri(data.trim());
      setFormState(FormState.SUCCESS);
      console.log(data);
    } catch (e) {
      setFormState(FormState.ERROR);
      setError(e);
      console.error(e);
    }
  };

  return (
    <>
      <div>
        <form onSubmit={handleSubmit}>
          <p className="text-center font-semibold text-3xl p-5">
            Upload the file here
          </p>
          <div>
            <input
              className="file:btn-form file:bg-red-400 file:hover:bg-red-500 file:shadow-red-300 w-full "
              type="file"
              id="file"
              onChange={(e) => setFile(e.target.files[0])}
            />
          </div>
          <div className="p-2">
            <button
              disabled={formState === FormState.SUBMITTING}
              className="btn-form flex flex-col bg-purple-600 hover:bg-purple-700 w-full py-5 hover:disabled:bg-slate-400 disabled:bg-slate-400"
            >
              {formState === FormState.SUBMITTING && (
                <div className="p-2 mx-auto">
                  <svg
                    version="1.1"
                    id="L3"
                    xmlns="http://www.w3.org/2000/svg"
                    xmlnsXlink="http://www.w3.org/1999/xlink"
                    x="0px"
                    y="0px"
                    viewBox="0 0 100 100"
                    enableBackground="new 0 0 0 0"
                    xmlSpace="preserve"
                    className="h-10"
                  >
                    <circle
                      fill="none"
                      stroke="#fff"
                      strokeWidth="4"
                      cx="50"
                      cy="50"
                      r="44"
                      style={{ opacity: "0.5" }}
                    />
                    <circle
                      fill="#fff"
                      stroke="#e74c3c"
                      strokeWidth="3"
                      cx="8"
                      cy="54"
                      r="6"
                    >
                      <animateTransform
                        attributeName="transform"
                        dur="2s"
                        type="rotate"
                        from="0 50 48"
                        to="360 50 52"
                        repeatCount="indefinite"
                      />
                    </circle>
                  </svg>
                </div>
              )}
              <p className="mx-auto">
                {formState === FormState.SUBMITTING ? "Processing..." : "OCR!"}
              </p>
            </button>
          </div>
        </form>

        <div>
          {error && <p className="text-red-500">{error}</p>}
          {formState === FormState.SUCCESS && downloadUri && (
            <a target="_blank" rel="noreferrer" href={downloadUri}>
              Download the file from here!
            </a>
          )}
        </div>
      </div>
    </>
  );
};
export default Form;
