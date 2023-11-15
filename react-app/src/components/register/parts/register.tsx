import { setFips } from 'crypto';
import React, { useEffect, useState, ChangeEvent, ChangeEventHandler } from 'react';
import { Link, NavigateFunction, useNavigate } from 'react-router-dom';

/**
 * データ登録
 * @returns 
 */
function Register() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [title, setTitle] = useState<string>('');
  const [tag, setTag] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const navigate: NavigateFunction = useNavigate()

  /**
   * タイトル入力値変更イベント
   * @param event
   */
  const handleTitleChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    // setTitle();
  };

  /**
   * タグ入力値変更イベント
   * @param event 
   */
  const handleTagChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
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
      const allowExtension: string[] = ["jpg", "jpeg", "png", "JPG", "JPEG", "PNG"];

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
          setMessage(data.error);
          return;
        }
        setMessage(data.message)

        // 入力値をリセットする
        setTitle("");
        setTag("");
        setSelectedFile(null);
        const thumbnailElem: HTMLInputElement = document.getElementById('thumbnail') as HTMLInputElement;
        thumbnailElem.value = '';

        navigate("/")
      } catch (error) {
        console.error('File upload failed', error);        
      }
    }
  };

  const resisterAreaCss ={
    maxWidth: '940px',
    padding: '0 40px',
    margin: '0 auto'
  }

  return (
    <div id='resister-area'>
        <div><p>{message}</p></div>
        <div>
            <textarea id='title' className='input-item' value={title} placeholder='Title' rows={1} onChange={handleTitleChange} ></textarea>
        </div>
        <div>
            <textarea id='tag' className='input-item' value={tag} placeholder='Tag' rows={1} onChange={handleTagChange} ></textarea>        
        </div>
        <div>
            <textarea id='article' className='input-item' value={tag} placeholder='write article content...' onChange={handleTagChange} ></textarea> 
        </div>
        <div>            
            <input id='thumbnail' className='input-item' type="file" accept=".jpg, .jpeg, .png" onChange={handleThumbnailChange} />            
        </div>
        <button onClick={handleUpload}>Upload</button> 
    </div>
  );
}

export default Register;
