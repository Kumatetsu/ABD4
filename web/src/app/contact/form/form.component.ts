import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import * as data from 'template/data.json';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';

@Component({
  selector: 'app-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.css']
})
export class FormComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english.contact;
  contactForm: FormGroup;
  submitted = false;
  title: string;
  name: string;
  email: string;
  message: string;
  nameError: string;
  emailError: string;
  emailFormatError: string;
  messageError: string;
  messageLengthError: string;
  submit: string;

  constructor(private formBuilder: FormBuilder, private translationService: TranslationService) {
    this.subscription = this.translationService.getLanguage().subscribe(language => {
      this.language = language.text;
      this.defineTemplate();
    });
  }
  ngOnInit() {
    this.contactForm = this.formBuilder.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      message: ['', [Validators.required, Validators.minLength(6)]]
    });
    this.language = this.translationService.getState().actual;
    this.defineTemplate();
  }
  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
  // return form controls
  get f() { return this.contactForm.controls; }
  // form submission callback functions
  onSubmit() {
    this.submitted = true;
    // stop here if form is invalid
    if (this.contactForm.invalid) {
      return;
    }
    alert(this.submit);
  }
  // variable setting relying on language
  defineTemplate() {
    if (this.language === this.translationService.getState().en) {
      this.template = (<any>data).english.contact;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese.contact;
    }
    this.nameError = this.template.nameError;
    this.title = this.template.title;
    this.name = this.template.name;
    this.email = this.template.email;
    this.message = this.template.message;
    this.emailError = this.template.emailError;
    this.emailFormatError = this.template.emailFormatError;
    this.messageError = this.template.messageError;
    this.messageLengthError = this.template.messageLengthError;
    this.submit = this.template.submit;
  }
}
