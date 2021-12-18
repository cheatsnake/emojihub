import { Test, TestingModule } from '@nestjs/testing';
import { EmojiController } from './emoji.controller';

describe('EmojiController', () => {
  let controller: EmojiController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [EmojiController],
    }).compile();

    controller = module.get<EmojiController>(EmojiController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
