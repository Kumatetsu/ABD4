import { Component, Input, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-tile',
  templateUrl: './tile.component.html',
  styleUrls: ['./tile.component.css']
})
export class TileComponent implements OnInit, OnDestroy {
  @Input() type: string;
  subscription: Subscription;
  language: string;
  template: any = (<any>data).english.tile.skills;
  tiles: any;
  title: string;
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
      this.template = (<any>data).english.tile.skills;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese.tile.skills;
    }
    this.tiles = this.template.tiles;
    this.title = this.template.title;
  }
}
