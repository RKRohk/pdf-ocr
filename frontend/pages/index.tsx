import Head from "next/head";
import Image from "next/image";
import { FormEvent, useState } from "react";
import Form, { FormState } from "../components/Form";
import Header from "../components/Header";

export default function Home() {
  return (
    <div className="p-5">
      <Head>
        <title>OCR PDF</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="h-full w-full">
        <Header />
        <div className="flex">
          <div className="mx-auto my-auto">
            <Form />
          </div>
        </div>
      </main>
    </div>
  );
}
