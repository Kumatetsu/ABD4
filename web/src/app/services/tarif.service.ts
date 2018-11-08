import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

interface Tarif {
    data: Object
}

@Injectable()
export class TarifService {
    url = "http://127.0.0.1:8000/";

    constructor (private http: HttpClient) { }

    getTarif(){
        return this.http.get<Tarif>(this.url + 'tarif')
    }

    postTarif(tarif) {
        let body = JSON.stringify(tarif);            
        let headers = new HttpHeaders({ 'Content-Type': 'application/json' });
            
        return this.http.post(this.url + 'tarif', body, { headers: headers })
    }
}
