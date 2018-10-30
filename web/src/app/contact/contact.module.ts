import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormComponent } from './form/form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { GmapComponent } from './gmap/gmap.component';
import { ContactComponent } from './contact/contact.component';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule
  ],
  declarations: [FormComponent, GmapComponent, ContactComponent],
  exports: [ContactComponent]
})
export class ContactModule { }
