module.exports = {
  HTML:function(user, admin, ui, printer) {
    let head, menu, list_menu, list = ``;
    if (ui == 'main') {
      if (user != null) {
        if (admin == "true") {
          head = `복합기 관리 솔루션<br/>
          ${user} 관리자님 어서오세요.`;
          list_menu = `<div class="col-sm-4">
            <ul class="nav nav-pills flex-column">
            <li class="nav-item">
              <a class="nav-link active" href="/">관리 복합기</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/service">서비스 내역</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">회원 관리</a>
            </li>
            </ul>
            <hr class="d-sm-none">
          </div>`;
        }else {
          head = `복합기 관리 솔루션<br/>
          ${user}님의 사용 복합기 리스트입니다.`;
          list_menu = `<div class="col-sm-4">
            <ul class="nav nav-pills flex-column">
            <li class="nav-item">
              <a class="nav-link active" href="#">사용 복합기</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">계정</a>
            </li>
            </ul>
            <hr class="d-sm-none">
          </div>`;
        }
        menu = `<li class="nav-item">
          <a class="nav-link" href="/logout_process">로그아웃</a>
        </li>`;
        list = `<div class="col-sm-8">
        <table class="table table-hover">
        <thead>
        <tr>
            <th scope="col">PrinterID</th>
            <th scope="col">Black</th>
            <th scope="col">Magenta</th>
            <th scope="col">Cyan</th>
            <th scope="col">Yellow</th>
            <th scope="col">Drum</th>
            <th scope="col">UsingPaper</th>
            <th scope="col">UserID</th>
        </tr>
        </thead>
        <tbody>`;
        list = list + printer;
        list = list + `</tbody>
          </table>
        <button type="button" class="btn btn-primary" onClick="location.href='/printer/enroll'">복합기 등록</button>`
      }
      else {
        head = `복합기 관리 솔루션<br/>
        로그인이 필요합니다.`;
        menu = `<li class="nav-item">
          <a class="nav-link" href="/login">로그인</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/sign_up">가입</a>
        </li>`;
        list_menu = `<p></p>`;
      }
    }else if (ui == 'enroll') { //복합기 등록 화면
      if (admin=="true") {
        list = `<div class="col-sm-8">
        <form action="http://localhost:3000/printer/enroll_process_a" method="post">
        <table>
          <tr>
            <td><div class="col-sm-1"><p>Printer ID</p></div></td>
            <td><p><input type="text" name="ID" placeholder="ID"></p></td>
          </tr>
          <tr>
            <td></td>
            <td><p>
              <input type="submit" class="btn btn-primary" value="등록"/>
              <button type="button" class="btn" onClick="location.href='/'">취소</button>
            </p></td>
          </tr>
        </table>
        </form>
        </div>`;
      }else {
        list = `<div class="col-sm-8">
        <form action="http://localhost:3000/printer/enroll_process" method="post">
        <table>
          <tr>
            <td><div class="col-sm-1"><p>Printer ID</p></div></td>
            <td><p><input type="text" name="ID" placeholder="ID"></p></td>
          </tr>
          <tr>
            <td><div class="col-sm-1"><p>IP</p></div></td>
            <td><p><textarea name="IP" placeholder="IP"></textarea></p></td>
          </tr>
          <tr>
            <td></td>
            <td><p>
              <input type="submit" class="btn btn-primary" value="등록"/>
              <button type="button" class="btn" onClick="location.href='/'">취소</button>
            </p></td>
          </tr>
        </table>
        </form>
        </div>`;
      }
      head = `${user}
      복합기 등록`;
      menu = `<li class="nav-item">
      <a class="nav-link" href="/logout_process">로그아웃</a>
      </li>`;
      list_menu = `<p></p>`; 
    }else if (ui == 'signup') { //회원가입 화면
      head = "회원가입";
      menu = `<li class="nav-item">
      <a class="nav-link" href="/login">로그인</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="/sign_up">가입</a>
      </li>`;
      list_menu = `<p></p>`;
      list =`<div class="col-sm-8">
      <form action="http://localhost:3000/sign_process" method="post">
        <table>
          <tr>
            <td><div class="col-sm-1"><p>ID</p></div></td>
            <td><p><input type="text" name="ID" placeholder="ID"></p></td>
          </tr>
          <tr>
            <td><div class="col-sm-1"><p>PW</p></div></td>
            <td><p><input type="text" name="PW" placeholder="PW"></p></td>
          </tr>
          <tr>
          <td></td>
          <td><p>
            <input type="submit" class="btn btn-primary" value="등록"/>
            <button type="button" class="btn" onClick="location.href='/'">취소</button>
          </p></td>
        </table>
      </form>
      </div>`;
    }else if (ui == 'login') {
      head = "로그인";
      menu = `<li class="nav-item">
      <a class="nav-link" href="/login">로그인</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="/sign_up">가입</a>
      </li>`;
      list_menu = `<p></p>`;
      list =`<div class="col-sm-8">
      <form action="http://localhost:3000/login_process" method="post">
      <table>
        <tr>
          <td><div class="col-sm-1"><p>ID</p></div></td>
          <td><p><input type="text" name="ID" placeholder="ID"></p></td>
        </tr>
        <tr>
          <td><div class="col-sm-1"><p>PW</p></div></td>
          <td><p><input type="text" name="PW" placeholder="PW"></p></td>
        </tr>
        <tr>
          <td></td>
          <td><p>
            <input type="submit" class="btn btn-primary" value="로그인"/>
            <button type="button" class="btn" onClick="location.href='/'">취소</button>
          </p></td>
        </tr>
      </table>
      </form>
      </div>`;
    }
    return UI(head, menu, list_menu, list);
  },list:function(id ,queryObj,user){
    var list = ``;
    list = list + `<tr>`;
    list = list + `<td>${id}</td>`
    list = list + `<td>${queryObj.black}</td>`
    list = list + `<td>${queryObj.magenta}</td>`
    list = list + `<td>${queryObj.cyan}</td>`
    list = list + `<td>${queryObj.yellow}</td>`
    list = list + `<td>${queryObj.drum}</td>`
    list = list + `<td>${queryObj.Paper}</td>`
    list = list + `<td>${user}</td>`
    list = list + `</tr>`;
    return list;
  }
}

function UI(head, menu, list_menu, list) {
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
    <a class="navbar-brand" href="/">복합기</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="collapsibleNavbar">
      <ul class="navbar-nav">
        ${menu}
      </ul>
    </div>
  </nav>

  <div class="container" style="margin-top:30px">
    <div class="row">
      ${list_menu}
      ${list}
      </div>
    </div>
  </div>
  </body>
  </html>
  `;
}