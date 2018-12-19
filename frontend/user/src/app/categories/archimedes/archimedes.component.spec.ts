import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ArchimedesComponent } from './archimedes.component';

describe('ArchimedesComponent', () => {
  let component: ArchimedesComponent;
  let fixture: ComponentFixture<ArchimedesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ArchimedesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ArchimedesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
