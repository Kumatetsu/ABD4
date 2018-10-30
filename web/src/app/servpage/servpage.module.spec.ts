import { ServpageModule } from './servpage.module';

describe('ServpageModule', () => {
  let servpageModule: ServpageModule;

  beforeEach(() => {
    servpageModule = new ServpageModule();
  });

  it('should create an instance', () => {
    expect(servpageModule).toBeTruthy();
  });
});
