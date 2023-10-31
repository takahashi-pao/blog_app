import { setFips } from 'crypto';
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
      // ファイルが選択された場合、拡張をチェックする
      const fileName: string = files[0].name;
      const fileExtension: string = fileName.split(".")[1];
      const allowExtension: string[] = ["jpg", "jpeg", "png"];

      if (allowExtension.includes(fileExtension) && fileName.split(".").length == 2){
        setSelectedFile(files[0]);
        return
      }

      // 不正な拡張子のファイルが選択された場合、選択を無効化する
      setMessage("不正な拡張子です");
      setSelectedFile(null);
      event.target.value = "";      
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
        const response = await fetch('http://localhost:8080/auth/Register', {
          method: 'POST',
          body: formData,
        });       

        const data = await response.json();
        if (data.error != null) {
          setMessage(data.error)
          return
        }
        setMessage(data.message)
        setTitle("")
        setTag("")
        setSelectedFile(null)
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
            <input type="file" accept=".jpg, .jpeg, .png" onChange={handleThumbnailChange} />            
        </div>
        <button onClick={handleUpload}>Upload</button>  
    </div>
  );
}

export default Register;
