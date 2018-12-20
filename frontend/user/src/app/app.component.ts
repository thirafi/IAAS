import { Component } from '@angular/core';
import { Http, Headers , RequestOptions,HttpModule } from "@angular/http";
import {HttpClientModule, HttpClient} from '@angular/common/http';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(
    public http: Http,
    public httpclientmodule :HttpClientModule,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    
   }
  title = 'Expronas Dashboard';
  isValid(): boolean {
    if ((this.router.url != '/Login') && (this.router.url != '/Register')) {
              return true;
      }
    return false;
  }
}
