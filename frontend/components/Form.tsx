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

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData();

    formData.append("file", file); //TODO (Fix unsafe type conversion)

    try {
      setFormState(FormState.SUBMITTING);
      const response = await fetch("http://localhost:8080/ocr", {
        method: "post",
        body: formData,
      });
      const data = response.url;
      setDownUri(data.trim());
      setFormState(FormState.SUCCESS);
      console.log(data);
    } catch (e) {
      setFormState(FormState.ERROR);
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
              className="file:btn-form file:bg-red-400 file:hover:bg-red-500 file:shadow-red-300"
              type="file"
              id="file"
              onChange={(e) => setFile(e.target.files[0])}
            />
          </div>
          <div className="p-2">
            <button
              disabled={formState === FormState.SUBMITTING}
              className="btn-form flex flex-col bg-purple-500 hover:bg-gradient-to-br w-full hover:from-purple-400 hover:to-purple-600 hover:scale-110 transition-transform duration-150 py-5 disabled:bg-slate-400"
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
                    enable-background="new 0 0 0 0"
                    xmlSpace="preserve"
                    className="h-10"
                  >
                    <circle
                      fill="none"
                      stroke="#fff"
                      stroke-width="4"
                      cx="50"
                      cy="50"
                      r="44"
                      style={{ opacity: "0.5" }}
                    />
                    <circle
                      fill="#fff"
                      stroke="#e74c3c"
                      stroke-width="3"
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
          {downloadUri && (
            <a target="_blank" href={downloadUri}>
              Download the file from here!
            </a>
          )}
        </div>
      </div>
    </>
  );
};
export default Form;
