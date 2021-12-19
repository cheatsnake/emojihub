import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { EmojiRandomService } from "./emoji-random.service";
import { EmojiController } from "./emoji.controller";
import { EmojiService } from "./emoji.service";
import { Emoji, EmojiSchema } from "./schemas/emoji.schema";

@Module({
    controllers: [EmojiController],
    providers: [EmojiService, EmojiRandomService],
    imports: [
        MongooseModule.forFeature([{ name: Emoji.name, schema: EmojiSchema }]),
    ],
})
export class EmojiModule {}
