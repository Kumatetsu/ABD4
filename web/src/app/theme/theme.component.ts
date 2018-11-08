import { Component, OnInit } from '@angular/core';
import { ThemeService } from '../services/theme.service';
import { HttpClient } from '@angular/common/http';
import { NotifierService } from 'angular-notifier';

@Component({
  selector: 'app-theme',
  templateUrl: './theme.component.html',
  styleUrls: ['./theme.component.css']
})

export class ThemeComponent {
    themes: any;
    new_description: any = "";
    error_msg: boolean = false;

    constructor(private ThemeService: ThemeService) {
    }

    ngOnInit() {
        this.getAllTheme();
    }

    addNewTheme() {
        if (this.new_description == ""){
            console.log("Description empty");

            this.error_msg = true;
            return;
        }

        let new_theme = {
            "Theme" : this.new_description
        }

        this.ThemeService.postTheme(new_theme).subscribe(
            response => console.log(response),
            error => console.log("Error : ", error),
            () => {
              this.getAllTheme()
              this.new_description = ""
            }
          )
    }

    getAllTheme() {
        let load_theme = true;
        this.ThemeService.getAllTheme().subscribe(
            response => this.themes = response.data,
            error => console.log("Error : ", error),
            () => load_theme = false
        )
    }  
}
