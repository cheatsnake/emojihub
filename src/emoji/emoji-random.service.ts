import { Injectable } from "@nestjs/common";
import { EmojiService } from "./emoji.service";
import { Emoji } from "./schemas/emoji.schema";

@Injectable()
export class EmojiRandomService {
    constructor(private readonly emojiService: EmojiService) {}

    async getRandomEmoji() {
        const emojis: Emoji[] = await this.emojiService.getAllEmoji();
        const randomNumber = getRandomNumber(emojis.length - 1);
        return emojis[randomNumber];
    }

    async getRandomEmojiInCategory(category: string) {
        const emojis: Emoji[] = await this.emojiService.getEmojiByCategory(
            category
        );
        const randomNumber = getRandomNumber(emojis.length - 1);
        return emojis[randomNumber];
    }

    async getRandomEmojiInGroup(group: string) {
        const emojis: Emoji[] = await this.emojiService.getEmojiByGroup(group);
        const randomNumber = getRandomNumber(emojis.length - 1);
        return emojis[randomNumber];
    }
}

const getRandomNumber = (max: number, min = 0): number => {
    return Math.round(Math.random() * (max - min)) + min;
};
