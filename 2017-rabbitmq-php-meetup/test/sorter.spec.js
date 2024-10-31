const should = require('chai').should();

const Message = require('./../src/message');
const Sorter = require('./../src/sorter');

/**
 * RUN these tests separately via: $ node_modules/mocha/bin/_mocha test/sorter.spec.js
 */
describe('Sorter', () => {
  it('should return data in correct order when all inits before resolves', () => {
    const sorter = new Sorter();
    sorter.init(100);
    sorter.init(101);
    sorter.init(102);
    sorter.init(103);
    sorter.init(104);

    let res = sorter.getMessages(100, '100data');
    res.should.deep.equal(['100data']);

    res = sorter.getMessages(101, '101data');
    res.should.deep.equal(['101data']);

    res = sorter.getMessages(104, '104data');
    res.should.deep.equal([]);

    res = sorter.getMessages(103, '103data');
    res.should.deep.equal([]);

    res = sorter.getMessages(102, '102data');
    res.should.deep.equal(['102data', '103data', '104data']);
  });
  it('should return data in correct order with mixed inits and resolves', () => {
    const sorter = new Sorter();
    sorter.init(100);
    sorter.init(101);

    let res = sorter.getMessages(100, '100data');
    res.should.deep.equal(['100data']);

    sorter.init(102);
    sorter.init(103);

    res = sorter.getMessages(101, '101data');
    res.should.deep.equal(['101data']);

    sorter.getMessages(103, '103data');
    res = sorter.getMessages(102, '102data');
    res.should.deep.equal(['102data', '103data']);

    sorter.init(104);

    res = sorter.getMessages(104, '104data');
    res.should.deep.equal(['104data']);
  });
  it('should not fail when trying to resolve before init', () => {
    const sorter = new Sorter();
    sorter.getMessages('xyz', 'data');
    sorter.init('xyz');
    const res = sorter.getMessages('xyz', 'data-b');
    res.should.deep.equal(['data-b']);
  });
  it('should not modify the stored original message object', () => {
    const sorter = new Sorter();

    const msgX = new Message({ job_id: 'X-123', sorteruence_id: 1 }, 'content-X-123');
    const msgY = new Message({ job_id: 'Y-456', sorteruence_id: 2 }, 'content-Y-123');

    sorter.init(msgX.getId());
    sorter.init(msgY.getId());

    sorter.getMessages(msgY.getId(), msgY);
    const res = sorter.getMessages(msgX.getId(), msgX);

    res.should.deep.equal([msgX, msgY]);
  });
});
