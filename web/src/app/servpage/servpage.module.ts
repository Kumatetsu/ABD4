import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WhoisComponent } from './whois/whois.component';
import { SkillsComponent } from './skills/skills.component';
import { PartnersComponent } from './partners/partners.component';
import { ServpageComponent } from './servpage/servpage.component';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [WhoisComponent, SkillsComponent, PartnersComponent, ServpageComponent],
  exports: [ServpageComponent]
})
export class ServpageModule { }
