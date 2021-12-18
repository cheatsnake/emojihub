import { Test, TestingModule } from '@nestjs/testing';
import { EmojiService } from './emoji.service';

describe('EmojiService', () => {
  let service: EmojiService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [EmojiService],
    }).compile();

    service = module.get<EmojiService>(EmojiService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
