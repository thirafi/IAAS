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
import { CeresComponent } from './categories/ceres/ceres.component';
import { ArchimedesComponent } from './categories/archimedes/archimedes.component';
import { LibertasComponent } from './categories/libertas/libertas.component';
import { LogbookComponent } from './pages/logbook/logbook.component';
import { StatusComponent } from './pages/status/status.component';

const appRoutes: Routes = [
  { path: '', redirectTo: '/Home', pathMatch:'full' },
  { path: 'Home', component: HomeComponent },
  { path: 'Register',  component: RegisterComponent },
  { path: 'Logbook', component: LogbookComponent },
  { path: 'Status', component: StatusComponent },
  { path: 'Login', component: LoginComponent },
  { path: 'HostingPlace', component: HostingplaceComponent },
  { path: 'Archimedes', component: ArchimedesComponent },
  { path: 'Libertas', component: LibertasComponent },
  { path: 'Ceres', component: CeresComponent },
  { path: 'Dashboard', component: DashboardComponent },


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
    LoginComponent,
    CeresComponent,
    ArchimedesComponent,
    LibertasComponent,
    LogbookComponent,
    StatusComponent
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
