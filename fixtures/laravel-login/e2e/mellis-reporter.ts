import type {
  FullConfig,
  FullResult,
  Reporter,
  TestCase,
  TestResult,
} from '@playwright/test/reporter';

function emit(status: string, name: string, extra = ''): void {
  const suffix = extra ? ` ${extra}` : '';
  console.log(`mellis:test ${status} ${name}${suffix}`);
}

class MellisReporter implements Reporter {
  onBegin(_config: FullConfig, suite: { allTests: () => TestCase[] }): void {
    for (const test of suite.allTests()) {
      emit('waiting', test.title);
    }
  }

  onTestBegin(test: TestCase): void {
    emit('running', test.title);
  }

  onTestEnd(test: TestCase, result: TestResult): void {
    const duration = `duration=${Math.round(result.duration)}ms`;
    if (result.status === 'passed') {
      emit('passed', test.title, duration);
      return;
    }

    const failure = (result.error?.message ?? result.status)
      .replace(/\s+/g, ' ')
      .trim();
    emit('failed', test.title, `${duration} failure=${failure}`);
  }

  onEnd(result: FullResult): void {
    void result;
  }
}

export default MellisReporter;
