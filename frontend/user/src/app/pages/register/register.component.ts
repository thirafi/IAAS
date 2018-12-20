import { Component, OnInit } from '@angular/core';
import { Http, Headers , RequestOptions,HttpModule } from "@angular/http";
import {HttpClientModule, HttpClient} from '@angular/common/http';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  accountData : { email?:string; role?:"user"; password?:string; }={};
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
    // this.accountData.email= "thirafi_aja";
    console.log("ini berhasil",  this.accountData );

  }
  register(){
    // let headers = new Headers({'Authorization':'Basic YWRtaW46YWRtaW4xMjM=','Content-Type':'application/json'});
    let headers = new Headers({'Content-Type':'application/json'});
    let options = new RequestOptions({headers: headers});
    // console.log(this.petugasData);
      console.log("ini berhasil",  this.accountData );
   
    this.http.post("http://localhost:8000/register",this.accountData,options).subscribe(data => {
      let response = data.json();
      console.log("ini berhasil", response );
      this.router.navigate(['/Login']);
    }, err => {     
      console.log("error nya apa: ",err);
    });
  }
}
