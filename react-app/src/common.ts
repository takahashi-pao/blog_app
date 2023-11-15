// 引数のバイト数を取得
export function GetByteSize(str: string): number {
    // Blobを利用してUTF-8エンコードした後のバイト数を取得
    const encoder = new TextEncoder();
    const byteArray = encoder.encode(str);

    // バイト数を返す
    return byteArray.length;
}