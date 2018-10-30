import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from './services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english;
  background: string;

  constructor(private translationService: TranslationService) {
    this.subscription = this.translationService.getLanguage().subscribe(language => {
      this.language = language.text;
      this.defineTemplate();
    });
  }
  ngOnInit() {
    this.language = this.translationService.getState().actual;
    this.defineTemplate();
  }
  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
  getBg() {
    return 'url(' + this.background + ') no-repeat center';
  }
  defineTemplate() {
    if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese;
    } else if (this.language === this.translationService.getState().en) {
      this.template = (<any>data).english;
    }
    this.background = this.template.background;
  }
}
