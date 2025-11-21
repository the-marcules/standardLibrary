'use client'
import React, { ChangeEvent, useEffect } from 'react'

export default function UploadPage() {
  const [file, setFile] = React.useState<File | null>(null)
  const [base64String, setBase64String] = React.useState<string>('')

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setFile(e.target.files[0])
    }
  }

  useEffect(() => {
    const encodeFileToBase64 = async (file: File) => {
      console.log('file selected:', file?.name)
      try {
        console.log('array buffer', await file.arrayBuffer())
        console.log('arrau buffer to string', (await file.arrayBuffer()).toString())
        const base64 = btoa((await file.arrayBuffer()).toString())
        setBase64String(base64)
        console.log('Base64 string:', base64)
      } catch (e) {
        console.error('Error converting file to base64:', e)
      }
    }
    if (file) {
      encodeFileToBase64(file!)
    }
  }, [file])

  return (
    <div>
      <h1>Upload Page</h1>
      <div>
        <label>
          Upload file:
          <input type="file" name="file" id="file" onChange={(e) => handleFileChange(e)} />
        </label>
      </div>
    </div>
  )
}
