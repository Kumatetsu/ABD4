import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

interface Theme {
    data: Object
}

@Injectable()
export class ThemeService {
    url = "http://127.0.0.1:8000/";

    constructor (private http: HttpClient) { }

    getAllTheme(){
        return this.http.get<Theme>(this.url + 'theme')
    }

    postTheme(theme) {
        let body = JSON.stringify(theme);            
        let headers = new HttpHeaders({ 'Content-Type': 'application/json' });
            
        return this.http.post(this.url + 'theme', body, { headers: headers })
    }
}
