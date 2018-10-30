import { Component, OnInit, OnDestroy } from '@angular/core';
import * as data from 'template/data.json';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';

@Component({
  selector: 'app-skills',
  templateUrl: './skills.component.html',
  styleUrls: ['./skills.component.css']
})
export class SkillsComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english.skills;
  title: string;
  phrases: any;
  tiles: any;
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
      this.template = (<any>data).english.skills;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese.skills;
    }
    this.tiles = this.template.tiles;
    this.title = this.template.title;
  }
}
