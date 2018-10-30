import { Component, OnInit, OnDestroy } from '@angular/core';
import * as data from 'template/data.json';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';

@Component({
  selector: 'app-servpage',
  templateUrl: './servpage.component.html',
  styleUrls: ['./servpage.component.css']
})
export class ServpageComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english;
  servimage = this.template.servimage;
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
      this.template = (<any>data).english;
    }
    this.servimage = this.template.servimage;
  }
}
