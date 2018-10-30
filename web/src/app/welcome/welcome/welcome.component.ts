import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  constructor(private translationService: TranslationService) {
    this.subscription = this.translationService.getLanguage().subscribe(language => this.language = language.text);
  }
  ngOnInit() {
    this.language = this.translationService.getState().actual;
  }
  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
}
