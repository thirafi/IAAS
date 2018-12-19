import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HostingplaceComponent } from './hostingplace.component';

describe('HostingplaceComponent', () => {
  let component: HostingplaceComponent;
  let fixture: ComponentFixture<HostingplaceComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HostingplaceComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HostingplaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
