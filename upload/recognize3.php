<?php
require_once("php_python.php"); //框架提供的程序脚本

// Support CORS
header("Access-Control-Allow-Origin: *");

/*
Windows bitmaps - *.bmp
JPEG 文件 - .jpeg, .jpg, *.jpe
JPEG 2000 文件 - *.jp2
Portable Network Graphics - *.png
*/
$fileType = array('jpg','jpe','jpeg', 'bmp', 'tif');
$max_file_size=2097152; 

$downloadDir = 'upload';
// 创建目录
if (!file_exists($downloadDir)) {
    @mkdir($downloadDir);
}

if ($_SERVER['REQUEST_METHOD'] == 'POST'){
        $file = $_FILES["file"];

        // $ret = array(
        // "errcode"=>0,
        // "errmsg"=>'ok',
        // "result"=>array()
        // );

        $headers = getallheaders();
        if (array_key_exists("X-Api-Key",$headers)) 
        { 
            if($headers['X-Api-Key'] == 'CFE95B64AC715D64275365EDE690GH7C')
            {
                //echo "Key exists!";
            }else{
                // $ret['errcode'] = 10001;
                // $ret['errmsg'] = "invalid X-Api-Key.";
                // die(json_encode($ret));
                die('{"errcode": 10001, "errmsg": "invalid X-Api-Key."}');
            }
        }else
        {
            // $ret['errcode'] = 10001;
            // $ret['errmsg'] = "invalid X-Api-Key.";
            // die(json_encode($ret));
            die('{"errcode": 10001, "errmsg": "invalid X-Api-Key."}');
        }

         var_dump($file);
	    die();
        //是否存在文件
		if (!is_uploaded_file($file['tmp_name']))
		{
             die('{"error" : {"code": 101, "message": "the image does not exist."}}');
		}
		//检查文件大小
		if($max_file_size < $file["size"])
		{
			die('{"error" : {"code": 101, "message": "the image size exceeds configured limit."}}');
		}
        //检查文件类型
        if (!in_array(str_replace("image/","",strtolower($file["type"])), $fileType))
		{
			die('{"error" : {"code": 101, "message": "the type of image is not allowed."}}');          
		}
	
		$filename=$file["tmp_name"];
		$pinfo=pathinfo($file["name"]);
		$ftype=$pinfo['extension'];
		$downloadPath = $downloadDir . DIRECTORY_SEPARATOR .uniqid().".".$ftype;
	
		if(!move_uploaded_file ($filename, $downloadPath))
		{
			die('{"error" : {"code": 101, "message": "failed to move uploaded file."}}');
		}
        
	
}else{
        if(empty($_GET['url'])||trim($_GET['url'])==""){
            die('{"error" : {"code": 100, "message": "url should not be NULL."}}');
        }

        $url = $_GET['url'];

        // 获取文件原文件名
        $defaultFileName = basename($url);

        // 获取文件类型
        $suffix = substr(strrchr($url, '.'), 1);
        $tmp_arr = explode("@",$suffix);
        $suffix = $tmp_arr[0];

        if (!in_array(strtolower($suffix), $fileType))
        {
            die('{"error" : {"code": 101, "message": "the type of image is not allowed."}}');
        }
        
        $downloadPath = $downloadDir . DIRECTORY_SEPARATOR . uniqid(). '-' . $defaultFileName;

        if(!file_exists(download_image($url,$downloadPath)))
        {
            die('{"error" : {"code": 102, "message": "failed to download the image."}}');
        }

}

$ret = array();

//"ppython"是框架"php_python.php"提供的函数，用来调用Python端服务
//调用Python的predict函数，并传递1个参数。
$ret['result'] = ppython("predict",dirname(__FILE__). DIRECTORY_SEPARATOR .$downloadPath);

$ret['result']['1'] = $ret['result']['1']/100;

//返回预测结果
echo json_encode($ret);

/**
 * 下载远程图片到本地
 *
 * @param string $url 远程文件地址
 * @param string $downloadPath 保存后的文件名
 * @param int $type 远程获取文件的方式
 * @return 文件路径
 */
function download_image($url, $downloadPath, $type = 1)
{
    // 获取远程文件资源
    if ($type){
        $ch = curl_init();
        $timeout = 30;
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, $timeout);
        $file = curl_exec($ch);
        curl_close($ch);
    }else{
        ob_start();
        readfile($url);
        $file = ob_get_contents();
        ob_end_clean();
    }
 
    // 保存文件
    $res = fopen($downloadPath, 'a');
    fwrite($res, $file);
    fclose($res);
 
    return $downloadPath;
}

function getallheaders()   
{  
    foreach ($_SERVER as $name => $value)   
    {  
        if (substr($name, 0, 5) == 'HTTP_')   
        {  
            $headers[str_replace(' ', '-', ucwords(strtolower(str_replace('_', ' ', substr($name, 5)))))] = $value;  
        }  
    }  
    return $headers;  
}  
