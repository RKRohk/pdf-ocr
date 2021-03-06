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
        <meta
          name="description"
          content="A completely free online service that adds an invisible layer of text over a PDF of images"
        />
        <link
          rel="apple-touch-icon"
          sizes="57x57"
          href="/apple-icon-57x57.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="76x76"
          href="/apple-icon-76x76.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="60x60"
          href="/apple-icon-60x60.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="72x72"
          href="/apple-icon-72x72.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="114x114"
          href="/apple-icon-114x114.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="120x120"
          href="/apple-icon-120x120.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="144x144"
          href="/apple-icon-144x144.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="152x152"
          href="/apple-icon-152x152.png"
        />
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/apple-icon-180x180.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="192x192"
          href="/android-icon-192x192.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="96x96"
          href="/favicon-96x96.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon-16x16.png"
        />
        <link rel="manifest" href="/manifest.json" />
        <meta name="msapplication-TileColor" content="#ffffff" />
        <meta name="msapplication-TileImage" content="/ms-icon-144x144.png" />
        <meta name="theme-color" content="#ffffff" />
        <meta property="og:type" content="website" />
        <meta property="og:url" content="https://ocr.rohankakar.me" />
        <meta property="og:title" content="OCR PDF" />
        <meta
          property="og:description"
          content="A completely free online service that adds an invisible layer of text over a PDF of images"
        />
        <meta property="og:image" content="/logo.png" />
        <meta property="twitter:card" content="summary_large_image"></meta>
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
