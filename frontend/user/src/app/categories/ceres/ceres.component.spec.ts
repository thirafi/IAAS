import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CeresComponent } from './ceres.component';

describe('CeresComponent', () => {
  let component: CeresComponent;
  let fixture: ComponentFixture<CeresComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CeresComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CeresComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
