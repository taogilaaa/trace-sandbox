import { Span, FORMAT_HTTP_HEADERS, Tags } from 'opentracing';
import { tracer } from '../../generals/tracer';
import { Metadata } from 'grpc';
import { saleOrderServiceClient } from './grpc';
import {
  GetSaleOrderRequest,
  GetSaleOrderResponse,
  GetSaleOrdersRequest,
  GetSaleOrdersResponse,
} from '../../generated/sandbox_sales_v1_pb';

type GetSaleOrderArgs = {
  id: number;
  span: Span;
};

type GetSaleOrdersArgs = {
  email: string;
  span: Span;
};

async function getSaleOrder({ span, ...args }: GetSaleOrderArgs) {
  const cSpan = tracer.startSpan('repository.getSaleOrder', {
    childOf: span,
    tags: {
      [Tags.SPAN_KIND]: Tags.SPAN_KIND_RPC_CLIENT,
      [Tags.PEER_ADDRESS]: saleOrderServiceClient.getChannel().getTarget(),
    },
  });

  const req = new GetSaleOrderRequest();

  req.setId(args.id);

  // hardcoded 15 second, might want to put in env only when needed
  const timeoutMs = 15000;
  const options = { deadline: Date.now() + timeoutMs };
  const meta = new Metadata();

  // https://github.com/jaegertracing/jaeger-client-node/pull/381#issuecomment-528182565
  tracer.inject(
    cSpan,
    FORMAT_HTTP_HEADERS,
    new Proxy(meta, {
      set(target, key: string, value) {
        target.set(key, value);
        return true;
      },
    }),
  );

  return new Promise<GetSaleOrderResponse.AsObject>((resolve, reject) => {
    saleOrderServiceClient.getSaleOrder(req, meta, options, (err, resp) => {
      if (err || !resp) {
        cSpan.setTag(Tags.ERROR, true);
        cSpan.finish();

        reject(err);
        return;
      }

      cSpan.finish();
      resolve(resp.toObject());
    });
  });
}

async function getSaleOrders({ span, ...args }: GetSaleOrdersArgs) {
  const cSpan = tracer.startSpan('repository.getSaleOrders', {
    childOf: span,
    tags: {
      [Tags.SPAN_KIND]: Tags.SPAN_KIND_RPC_CLIENT,
      [Tags.PEER_ADDRESS]: saleOrderServiceClient.getChannel().getTarget(),
    },
  });

  const req = new GetSaleOrdersRequest();

  req.setEmail(args.email);

  // hardcoded 15 second, might want to put in env only when needed
  const timeoutMs = 15000;
  const options = { deadline: Date.now() + timeoutMs };
  const meta = new Metadata();

  // https://github.com/jaegertracing/jaeger-client-node/pull/381#issuecomment-528182565
  tracer.inject(
    cSpan,
    FORMAT_HTTP_HEADERS,
    new Proxy(meta, {
      set(target, key: string, value) {
        target.set(key, value);
        return true;
      },
    }),
  );

  return new Promise<GetSaleOrdersResponse.AsObject>((resolve, reject) => {
    saleOrderServiceClient.getSaleOrders(req, meta, options, (err, resp) => {
      if (err || !resp) {
        cSpan.setTag(Tags.ERROR, true);
        cSpan.finish();

        reject(err);
        return;
      }

      cSpan.finish();
      resolve(resp.toObject());
    });
  });
}

export { getSaleOrder, getSaleOrders };
