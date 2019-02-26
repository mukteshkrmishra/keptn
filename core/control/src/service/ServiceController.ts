import * as express from 'express';
import { inject, injectable } from 'inversify';
import {
  controller,
  httpGet,
  httpPost,
  interfaces,
  httpDelete,
} from 'inversify-express-utils';
import 'reflect-metadata';
import {
  ApiOperationGet,
  ApiOperationPost,
  ApiPath,
  SwaggerDefinitionConstant,
  ApiOperationDelete,
} from 'swagger-express-ts';

@ApiPath({
  name: 'Service',
  path: '/service',
  security: { apiKeyHeader: ['x-keptn-signature'] },
})
@controller('service')
export class ServiceController implements interfaces.Controller {

  @ApiOperationPost({
    description: 'Onboards a new service to a keptn project',
    parameters: {
      body: {
        description: 'Service information',
        model: 'ServiceRequestModel',
        required: true,
      },
    },
    responses: {
      200: {
      },
      400: { description: 'Parameters fail' },
    },
    summary: 'Onboard a new service to a keptn project',
  })
  @httpPost('/')
  public onboardService(
    request: express.Request,
    response: express.Response,
    next: express.NextFunction,
  ): void {
    const result = {
      result: 'success',
    };

    response.send(result);
  }

  @ApiOperationDelete({
    description: 'Delete a project',
    parameters: {

    },
    responses: {
      200: {
      },
      400: { description: 'Parameters fail' },
    },
    summary: 'Delete a keptn project',
  })
  @httpDelete('/')
  public deleteService(
    request: express.Request,
    response: express.Response,
    next: express.NextFunction,
  ): void {
    const result = {
      result: 'success',
    };

    response.send(result);
  }
}
