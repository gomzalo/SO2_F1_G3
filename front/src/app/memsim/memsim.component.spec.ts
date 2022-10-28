import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MemsimComponent } from './memsim.component';

describe('MemsimComponent', () => {
  let component: MemsimComponent;
  let fixture: ComponentFixture<MemsimComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MemsimComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MemsimComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
