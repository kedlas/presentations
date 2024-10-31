
const assert = require('assert');
const uuid = require('uuid/v1');

class Message {

  /**
   *
   * @param headers {Object}
   * @param content {String}
   */
  constructor(headers, content) {
    assert(headers.job_id, 'Missing \'job_id\' message header.');

    this.jobId = `${headers.job_id}`;
    this.id = `${headers.job_id}:${uuid()}`;
    this.content = content;
  }

  /**
   *
   * @return {string}
   */
  getId() {
    return this.id;
  }

  /**
   *
   * @return {string}
   */
  getJobId() {
    return this.jobId;
  }

  /**
   *
   * @return {string}
   */
  getContent() {
    return this.content;
  }

}

module.exports = Message;
