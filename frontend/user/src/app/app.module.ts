import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { HeaderComponent } from './elements/header/header.component';
import { SidebarComponent } from './elements/sidebar/sidebar.component';
import { FooterComponent } from './elements/footer/footer.component';
import { RightbarComponent } from './elements/rightbar/rightbar.component';
import { HomeComponent } from './pages/home/home.component';
import { RegisterComponent } from './pages/register/register.component';
import { HostingplaceComponent } from './pages/hostingplace/hostingplace.component';
import { AppRoutingModule } from './/app-routing.module';
import { LoginComponent } from './pages/login/login.component';

const appRoutes: Routes = [
  { path: '', redirectTo: '/Home', pathMatch:'full' },
  { path: 'Home', component: HomeComponent },
  { path: 'Register',      component: RegisterComponent },
  { path: 'Login', component: LoginComponent },
  { path: 'HostingPlace', component: HostingplaceComponent },

];

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    HeaderComponent,
    SidebarComponent,
    FooterComponent,
    RightbarComponent,
    HomeComponent,
    RegisterComponent,
    HostingplaceComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(
      appRoutes,
    ),
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
