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

const appRoutes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'Regsiter',      component: RegisterComponent },
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
    HostingplaceComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(
      appRoutes,
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
