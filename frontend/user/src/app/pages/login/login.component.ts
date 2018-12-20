import { Component, OnInit } from '@angular/core';
import { Http, Headers , RequestOptions,HttpModule } from "@angular/http";
import {HttpClientModule, HttpClient} from '@angular/common/http';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  submitted: boolean;
  accountData : { email?:string; password?:string; }={};
  constructor(
    public http: Http,
    public httpclientmodule :HttpClientModule,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    if (localStorage.getItem("token") != null) {
      this.router.navigate(['/Home']);
    }
   }

  ngOnInit() {
  }
  login(){
    this.submitted = true;
  
      let headers = new Headers({'Content-Type':'application/json'});
      let options = new RequestOptions({headers: headers});
      console.log("header :",options);
      this.http.post("http://localhost:8000/login",this.accountData,options).subscribe(data => {
        let response = data.json();
        console.log("sukses",response);
        localStorage.setItem("token", JSON.stringify(response.token));
        this.router.navigate(['/Home']);
      }, err => {     
        console.log("Anda Belum di daftarkan oleh Admin",err);
        
      });
    }
}
