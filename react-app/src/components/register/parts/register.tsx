import React, { useState, ChangeEvent } from 'react';

/**
 * データ登録
 * @returns 
 */
function Register() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [title, setTitle] = useState<string>('');
  const [tag, setTag] = useState<string>('');
  const [message, setMessage] = useState<string>('');

  /**
   * タイトル入力値変更イベント
   * @param event
   */
  const handleTitleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  /**
   * タグ入力値変更イベント
   * @param event 
   */
  const handleTagChange = (event: ChangeEvent<HTMLInputElement>) => {
    setTag(event.target.value);
  };

  /**
   * ファイル変更イベント
   * @param e 
   */
  const handleThumbnailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      setSelectedFile(files[0]);
    }
  };

  /**
   * 登録ボタン押下イベント
   */
  const handleUpload = async () => {
    if (selectedFile) {
      const formData = new FormData();
      formData.append('title', title);
      formData.append('tag', tag);
      formData.append('file', selectedFile);

      try {
        const response = await fetch('http://localhost:8080/Register', {
          method: 'POST',
          body: formData,
        });       

        const data = await response.json();
        setMessage(data.message)
      } catch (error) {
        console.error('File upload failed', error);        
      }
    }
  };

  return (
    <div>
        <div><p>{message}</p></div>
        <div>
            <label>Title</label>
            <input type='text' value={title} onChange={handleTitleChange} />
        </div>
        <div>
            <label>Tag</label>
            <input id='tag' value={tag} onChange={handleTagChange} />            
        </div>
        <div>
            <label>Thumbnail</label>
            <input type="file" accept="image/*" onChange={handleThumbnailChange} />            
        </div>
        <button onClick={handleUpload}>Upload</button>  
    </div>
  );
}

export default Register;
