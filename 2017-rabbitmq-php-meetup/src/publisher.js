'use strict';

const amqp = require('amqplib/callback_api');
const conf = require('./../config');

function randomInt(low, high) {
  return Math.floor((Math.random() * (high - low)) + low);
}

amqp.connect(conf.rabbit.url, (connErr, conn) => {
  if(connErr) {
    console.error(connErr);
    process.exit(0);
  }

  conn.createChannel((chanErr, ch) => {
    if(chanErr) {
      console.error(chanErr);
      process.exit(0);
    }

    ch.assertExchange(conf.rabbit.exchange.name, conf.rabbit.exchange.type, {}, (exErr) => {
      if (exErr) {
        console.error(exErr);
        process.exit(0);
      }

      let i = 1;
      while (i <= conf.message_count) {
        const random = randomInt(10, 10000);
        const content = `Message ${i}. My processing should take ${random}ms`;

        ch.publish(conf.rabbit.exchange.name, '', new Buffer(content), {
          messageId: `${i}`,
          headers: {
            job_id: 'test',
            sequence_id: `${i}`,
            duration: random
          }
        });
        console.info(`[x] Published ${i}. Content: "${content}"`);
        i += 1;
      }

      setTimeout(() => {
        conn.close();
        process.exit(0);
      }, 200);

    });
  });
});
