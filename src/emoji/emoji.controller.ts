import { Controller, Get } from "@nestjs/common";
import { EmojiService } from "./emoji.service";

@Controller("api")
export class EmojiController {
    constructor(private readonly emojiService: EmojiService) {}

    @Get("emoji")
    getAll() {
        return this.emojiService.getAll();
    }
}
