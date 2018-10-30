import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class TranslationService {
    private actual = 'English';
    private language = new Subject<any>();

    sendLanguage(language: string) {
        this.actual = language;
        this.language.next({ text: language });
    }

    clearQueue() {
        this.language.next();
    }

    getLanguage(): Observable<any> {
        return this.language.asObservable();
    }

    getState() {
        return {
            'en': 'English',
            'ch': '中文',
            'actual': this.actual,
        };
    }
}
