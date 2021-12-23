import Head from "next/head";
import Image from "next/image";
import { FormEvent, useState } from "react";
import Form, { FormState } from "../components/Form";

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
            <Form />
          </div>
        </div>
      </main>
    </div>
  );
}
