var express = require('express');
var template = require('./template');
var queryB = require('./queryB');
//var registerUser = require('./registerUser');
var bodyParser = require('body-parser');
var mysql = require('mysql2/promise');
var app = express();
app.use(bodyParser.urlencoded({extended:true}))
app.use(bodyParser.json())


app.get('/', async function(request, response) {
  //if (request.cookies) {
  //}
    var list = `<table class="table table-hover">
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
    // let obj = {
    // id : 'Epson cx-29',
    // black : 100,
    // magenta : 20,
    // cyan : 30,
    // yellow : 40,
    // drum : 20,
    // Paper : 24884
    // };
    // list = list + template.list(obj);

    let connection = await mysql.createConnection({
      host: 'localhost',
      user: 'nodejs',
      password: '1234',
      database: 'my_db',
    });

    let objQ = await connection.query("SELECT `printer` FROM `plist` where `id`=?", ["dori2005"])

    try {
      for (let i = 0; i < objQ.length; i++) {
        let queryjson = await queryB('updatePrint', objQ[0][i].printer, null);
        console.log(queryjson);
        let obj = JSON.parse(queryjson);
        list = list + template.list(objQ[0][i].printer, obj);
      }
    } catch (error) {
      console.log(error);
    }
    list = list + `  </tbody>
    </table>
   <button type="button" class="btn btn-primary" onClick="location.href='/enroll'">복합기 등록</button>
   <button type="button" class="btn">삭제</button>`

      var html = template.HTML(`복합기 관리 솔루션<br/>
      dori2005님의 사용 복합기 리스트입니다.`,list);

      response.send(html);
});

app.get('/sign_up', function(request, response) {
  var html = template.HTML("회원가입",`
              <form action="http://localhost:3000/sign_process" method="post">
                <p><input type="text" name="ID" placeholder="ID"></p>
                <p><input type="text" name="PW" placeholder="PW"></p>
                <p><input type="text" name="PID" placeholder="Printer ID"></p>
                <p>
                  <input type="submit" class="btn btn-primary" value="등록"/>
                  <button type="button" class="btn" onClick="location.href='/'">취소</button>
                </p>
              </form>
            `);
    response.send(html);
});

app.get('/login', function(request, response) {
    var html = template.HTML("로그인",`
                <form action="http://localhost:3000/login_process" method="post">
                  <p><input type="text" name="ID" placeholder="ID"></p>
                  <p><input type="text" name="PW" placeholder="PW"></p>
                  <p>
                    <input type="submit" class="btn btn-primary" value="로그인"/>
                    <button type="button" class="btn" onClick="location.href='/login_process'">취소</button>
                  </p>
                </form>
              `);
      response.send(html);
});

app.get('/enroll', function(request, response) {
    var html = template.HTML(`
              <form action="http://localhost:3000/enroll_process" method="post">
                <p><input type="text" name="ID" placeholder="ID"></p>
                <p>
                  <textarea name="IP" placeholder="IP"></textarea>
                </p>
                <p>
                  <input type="submit" class="btn btn-primary" value="등록"/>
                  <button type="button" class="btn" onClick="location.href='/'">취소</button>
                </p>
              </form>
            `);
    response.send(html);
});

app.post('/enroll_process', async function(request, response) {
    var id = request.body.ID;
    var ip = request.body.IP;
    
    pool.query("INSERT INTO `plist` VALUE(?,?,?)",["dori2005","1234",id], function(err, rows, fields){
      console.log(rows);
      // 쿼리가 수행되면 connection은 자동으로 해제된다.
    })
    let queryjson = await queryB('enrollPrint', id, ip);

    response.redirect('/')
});

app.post('/login_process', function(request, response) {
  var id = request.body.ID;
  var pw = request.body.PW;
  
  console.log(request.body)
  console.log(id)

  pool.query("INSERT INTO `plist` VALUE(?,?,?)",[id,pw,""], function(err, rows, fields){
    console.log(rows);
    // 쿼리가 수행되면 connection은 자동으로 해제된다.
  })

  pool.query("SELECT * FROM `plist`", function(err, rows, fields){
    console.log(rows);
    // 쿼리가 수행되면 connection은 자동으로 해제된다.
  })

  response.redirect('/')
});

app.post('/sign_process', function(request, response) {
    var id = request.body.ID;
    var pw = request.body.PW;
    var pid = request.body.PID;
    
    console.log(request.body)
    console.log(id)

    pool.query("INSERT INTO `plist` VALUE(?,?,?)",[id,pw,pid], function(err, rows, fields){
      console.log(rows);
      // 쿼리가 수행되면 connection은 자동으로 해제된다.
    })

    pool.query("SELECT * FROM `plist`", function(err, rows, fields){
      console.log(rows);
      // 쿼리가 수행되면 connection은 자동으로 해제된다.
    })

    response.redirect('/')
});

app.get('/test', function(request, response) {
    response.send('test');
});

app.listen(3000, function() {
    console.log('Example app listening on port 3000!');
})
