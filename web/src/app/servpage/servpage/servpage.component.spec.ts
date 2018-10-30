import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServpageComponent } from './servpage.component';

describe('ServpageComponent', () => {
  let component: ServpageComponent;
  let fixture: ComponentFixture<ServpageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServpageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServpageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
