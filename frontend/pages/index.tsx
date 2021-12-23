import Head from "next/head";
import Image from "next/image";
import Form from "../components/Form";

export default function Home() {
  return (
    <div className="p-5">
      <Head>
        <title>OCR PDF</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="h-full w-full">
        <h1>PDF OCR</h1>
        <div className="flex">
          <div className="mx-auto my-auto">
            <Form
              onSubmit={async (e) => {
                e.preventDefault();
                await new Promise((resolve) => setTimeout(resolve, 1000));
              }}
            />
          </div>
        </div>
      </main>
    </div>
  );
}
