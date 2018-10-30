import { WebbaseModule } from './webbase.module';

describe('WebbaseModule', () => {
  let webbaseModule: WebbaseModule;

  beforeEach(() => {
    webbaseModule = new WebbaseModule();
  });

  it('should create an instance', () => {
    expect(webbaseModule).toBeTruthy();
  });
});
