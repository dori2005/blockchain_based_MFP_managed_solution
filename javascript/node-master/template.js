module.exports = {
  HTML:function(head,list) {
    return `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <title>Capstone Design</title>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js"></script>
    </head>
    <body>

    <div class="jumbotron text-center" style="margin-bottom:0">
      <h1>${head}</h1>
    </div>

    <nav class="navbar navbar-expand-sm bg-secondary navbar-dark">
      <a class="navbar-brand" href="#">복합기</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="collapsibleNavbar">
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link" href="/login">로그인</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/sign_up">가입</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/enroll">로그아웃</a>
          </li>
        </ul>
      </div>
    </nav>

    <div class="container" style="margin-top:30px">
      <div class="row">
        <div class="col-sm-4">
          <ul class="nav nav-pills flex-column">
            <li class="nav-item">
              <a class="nav-link active" href="#">사용 복합기</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">서비스 내역</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">계정</a>
            </li>
          </ul>
          <hr class="d-sm-none">
        </div>
        <div class="col-sm-8">
        ${list}
        </div>
      </body>
    </html>
    `;
  },list:function(id ,queryObj){
    var list = ``;
    list = list + `<tr>`;
    list = list + `<td>${id}</td>`
    list = list + `<td>${queryObj.black}</td>`
    list = list + `<td>${queryObj.magenta}</td>`
    list = list + `<td>${queryObj.cyan}</td>`
    list = list + `<td>${queryObj.yellow}</td>`
    list = list + `<td>${queryObj.drum}</td>`
    list = list + `<td>${queryObj.Paper}</td>`
    list = list + `<td>"dori2005"</td>`
    list = list + `</tr>`;
    return list;
  }
}
