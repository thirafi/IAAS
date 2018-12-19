import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LibertasComponent } from './libertas.component';

describe('LibertasComponent', () => {
  let component: LibertasComponent;
  let fixture: ComponentFixture<LibertasComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LibertasComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LibertasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
