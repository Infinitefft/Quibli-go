import { IsString, IsArray, IsOptional, IsNotEmpty } from 'class-validator';

export class CreatePostDto {
  @IsString()
  @IsNotEmpty()
  title: string;

  @IsString()
  @IsNotEmpty()
  content: string;

  @IsArray()
  @IsOptional() // 如果 tags 可以不传，就加这个
  tags?: string[];
}