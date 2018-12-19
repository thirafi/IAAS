import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HostingplaceeeComponent } from './hostingplaceee.component';

describe('HostingplaceeeComponent', () => {
  let component: HostingplaceeeComponent;
  let fixture: ComponentFixture<HostingplaceeeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HostingplaceeeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HostingplaceeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
