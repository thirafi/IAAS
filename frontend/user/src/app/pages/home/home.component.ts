import { Component, OnInit } from '@angular/core';
import { Http, Headers , RequestOptions,HttpModule } from "@angular/http";
import {HttpClientModule, HttpClient} from '@angular/common/http';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  notif :boolean ;
  constructor(
    public http: Http,
    public httpclientmodule :HttpClientModule,
    private route: ActivatedRoute,
  private router: Router,
  ) {
  this.notif =true ;
    if (localStorage.getItem("token") == null) {
    this.router.navigate(['/Login']);} 
}

  ngOnInit() {
    
  }
  
}
