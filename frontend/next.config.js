module.exports = {
  reactStrictMode: true,
  async rewrites() {

    return [
      {
        source:"/api/ocr",
        destination:"http://localhost:8080/ocr"
      },
      {
        source:"/ocr/:fileID*",
        destination:"http://localhost:8080/ocr/:fileID*"
      }
    ]
  }
}
