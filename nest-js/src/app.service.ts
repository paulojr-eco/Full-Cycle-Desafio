import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getInfo(): string {
    return 'This app consists in a gRPC client, where you can make requests to create and list products.';
  }
}
