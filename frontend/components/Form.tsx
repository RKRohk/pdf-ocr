type FormProps = {
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
};
const Form: React.VFC<FormProps> = ({ onSubmit }) => {
  return (
    <>
      <form onSubmit={onSubmit}>
        <p className="text-center font-semibold text-3xl p-5">
          Upload the file here{" "}
        </p>
        <div>
          <input
            className="file:btn-form file:bg-red-400 file:hover:bg-red-500 file:shadow-red-300"
            type="file"
            id="file"
          />
        </div>
        <div className="p-2">
          <button className="btn-form bg-purple-500 hover:bg-gradient-to-br w-full hover:from-purple-400 hover:to-purple-600 hover:scale-110 transition-transform duration-150 py-5">
            OCR!
          </button>
        </div>
      </form>
    </>
  );
};
export default Form;
