import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-intro',
  templateUrl: './intro.component.html',
  styleUrls: ['./intro.component.css']
})
export class IntroComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english;
  addr: string;
  company: string;
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
  defineTemplate() {
    if (this.language === this.translationService.getState().en) {
      this.template = (<any>data).english;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese;
    }
    this.addr = this.template.addr;
    this.company = this.template.company;
  }
}
