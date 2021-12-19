import { Controller, Get, Param } from "@nestjs/common";
import { EmojiRandomService } from "./emoji-random.service";
import { EmojiService } from "./emoji.service";

@Controller("api")
export class EmojiController {
    constructor(
        private readonly emojiService: EmojiService,
        private readonly emojiRandomServise: EmojiRandomService
    ) {}

    @Get("all")
    getAllEmoji() {
        return this.emojiService.getAllEmoji();
    }

    @Get("all/category_:category")
    getAllEmojiInCategory(@Param("category") category: string) {
        return this.emojiService.getEmojiByCategory(category);
    }

    @Get("all/group_:group")
    getAllEmojiInGroup(@Param("group") group: string) {
        return this.emojiService.getEmojiByGroup(group);
    }

    @Get("random")
    getRandomEmoji() {
        return this.emojiRandomServise.getRandomEmoji();
    }

    @Get("random/category_:category")
    getRandomEmojiByCategory(@Param("category") category: string) {
        return this.emojiRandomServise.getRandomEmojiInCategory(category);
    }

    @Get("random/group_:group")
    getRandomEmojiByGroupInCategory(@Param("group") group: string) {
        return this.emojiRandomServise.getRandomEmojiInGroup(group);
    }
}
