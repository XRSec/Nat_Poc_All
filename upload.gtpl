<html>
  <head>
    <title>Upload file</title>
    <style>
      .box{
        width: 90vh;
        margin: 30vh auto;

      }
      iframe {
        width: 84vh;
        height: auto;
        display: flex;
        align-items: center;
        height: 55px;
        border: 2px dashed #ff6700;
        border-radius: 2px;
        margin: 0 auto;
        font-weight: bold;
        background: rgb(255 255 255 / 64%);
        padding: 0 3vh;
      }
      .upload {
        display: flex;
        align-items: center;
        padding: 0 60px;
        height: 100px;
        border: 2px dashed #409eff;
        border-radius: 2px;
        margin: 35px 0;
      }
      input[type="file"] {
        color: #ff6700;
        margin: 0 auto;
        font-weight: bold;
        background: #ffffff70;
      }
      input[type="submit"] {
        display: block;
        background: #409eff;
        width: 160px;
        height: 35px;
        border: none;
        border-radius: 3px;
        text-align: center;
        color: #fff;
        margin: 0 auto;
        /* cursor: pointer; */
        font-weight: bold;
      }
      input[type="file"]:hover {
        opacity: 0.7;
      }
      input[type="submit"]:hover {
        opacity: 0.7;
      }
      body {
        background: url(https://www.bing.com/th?id=OHR.Autumn_ROW4111709953_1920x1080.jpg&rf=LaDigue_1920x1080.jpg);
        top: 0;
        left: 0;
        width: 97%;
        height: auto;
        z-index: 2;
        background-position: center;
        background-size: cover;
      }
    </style>
  </head>
  <body>
    <div class="box">
      <form target="form" enctype="multipart/form-data" action="/upload" method="post">
        <div class="upload">
          <input type="file" name="uploadFile"/>
          <input type="submit" value="upload"/>
        </div>
      </form>
      <iframe name="form" id="form"></iframe>
    </div>
  </body>
</html>
