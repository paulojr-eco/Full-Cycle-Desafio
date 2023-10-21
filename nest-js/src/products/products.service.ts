import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreateProductDto } from './dto/create-product.dto';
import { ClientGrpc } from '@nestjs/microservices';
import { Product, ProductClientGrpc } from './products.grpc';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class ProductsService implements OnModuleInit {
  private productGrpcService: ProductClientGrpc;

  constructor(
    @Inject('PRODUCT_PACKAGE')
    private productGrpcPackage: ClientGrpc,
  ) {}

  onModuleInit() {
    this.productGrpcService =
      this.productGrpcPackage.getService('ProductService');
  }

  async create(createProductDto: CreateProductDto): Promise<Product> {
    return await lastValueFrom(
      this.productGrpcService.createProduct({
        ...createProductDto,
      }),
    );
  }

  async findAll() {
    return await lastValueFrom(this.productGrpcService.findProducts({}));
  }
}
